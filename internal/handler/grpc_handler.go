package handler

import (
	"context"
	"errors"
	"net/url"

	"github.com/Sorrowful-free/short-url-service/api"
	"github.com/Sorrowful-free/short-url-service/internal/consts"
	"github.com/Sorrowful-free/short-url-service/internal/crypto"
	"github.com/Sorrowful-free/short-url-service/internal/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// GRPCHandler implements the gRPC ShortenerService server
type GRPCHandler struct {
	api.UnimplementedShortenerServiceServer
	baseURL         string
	urlService      service.ShortURLService
	userIDEncryptor crypto.UserIDEncryptor
}

// NewGRPCHandler creates a new gRPC handler instance
func NewGRPCHandler(baseURL string, urlService service.ShortURLService, userIDEncryptor crypto.UserIDEncryptor) *GRPCHandler {
	return &GRPCHandler{
		baseURL:         baseURL,
		urlService:      urlService,
		userIDEncryptor: userIDEncryptor,
	}
}

// getUserIDFromMetadata extracts user ID from gRPC metadata
func (h *GRPCHandler) getUserIDFromMetadata(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		// If no metadata, generate a new user ID
		userID, err := crypto.GenerateRandomSequenceString(consts.TestUserIDLength)
		if err != nil {
			return consts.FallbackUserID, nil
		}
		return userID, nil
	}

	authHeaders := md.Get("authorization")
	if len(authHeaders) == 0 || authHeaders[0] == "" {
		// If no authorization header, generate a new user ID
		userID, err := crypto.GenerateRandomSequenceString(consts.TestUserIDLength)
		if err != nil {
			return consts.FallbackUserID, nil
		}
		return userID, nil
	}

	// Decrypt the user ID from the authorization header
	userID, err := h.userIDEncryptor.Decrypt(authHeaders[0])
	if err != nil {
		// If decryption fails, generate a new user ID
		newUserID, genErr := crypto.GenerateRandomSequenceString(consts.TestUserIDLength)
		if genErr != nil {
			return consts.FallbackUserID, nil
		}
		return newUserID, nil
	}

	return userID, nil
}

// ShortenURL creates a short URL from the provided original URL
func (h *GRPCHandler) ShortenURL(ctx context.Context, req *api.URLShortenRequest) (*api.URLShortenResponse, error) {
	if req.Url == "" {
		return nil, status.Error(codes.InvalidArgument, "url is required")
	}

	userID, err := h.getUserIDFromMetadata(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get user ID")
	}

	dto, err := h.urlService.TryMakeShort(ctx, userID, req.Url)
	var originalURLConflictError *service.OriginalURLConflictServiceError
	if err != nil && !errors.As(err, &originalURLConflictError) {
		return nil, status.Error(codes.Internal, err.Error())
	}

	shortURL, err := url.JoinPath(h.baseURL, dto.ShortUID)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to construct short URL")
	}

	return &api.URLShortenResponse{
		Result: shortURL,
	}, nil
}

// ExpandURL retrieves the original URL for the given short URL ID
func (h *GRPCHandler) ExpandURL(ctx context.Context, req *api.URLExpandRequest) (*api.URLExpandResponse, error) {
	if req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}

	dto, err := h.urlService.TryMakeOriginal(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, "short URL not found")
	}

	if dto.IsDeleted {
		return nil, status.Error(codes.NotFound, "short URL is deleted")
	}

	return &api.URLExpandResponse{
		Result: dto.OriginalURL,
	}, nil
}

// ListUserURLs returns all URLs associated with the authenticated user
func (h *GRPCHandler) ListUserURLs(ctx context.Context, _ *emptypb.Empty) (*api.UserURLsResponse, error) {
	userID, err := h.getUserIDFromMetadata(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get user ID")
	}

	shortURLDTOs, err := h.urlService.GetUserUrls(ctx, userID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if len(shortURLDTOs) == 0 {
		return &api.UserURLsResponse{
			Url: []*api.URLData{},
		}, nil
	}

	urlDataList := make([]*api.URLData, 0, len(shortURLDTOs))
	for _, shortURLDTO := range shortURLDTOs {
		shortURL, err := url.JoinPath(h.baseURL, shortURLDTO.ShortUID)
		if err != nil {
			return nil, status.Error(codes.Internal, "failed to construct short URL")
		}

		urlDataList = append(urlDataList, &api.URLData{
			ShortUrl:    shortURL,
			OriginalUrl: shortURLDTO.OriginalURL,
		})
	}

	return &api.UserURLsResponse{
		Url: urlDataList,
	}, nil
}
