package wasm

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"

	postskeeper "github.com/warmage-sports/mage/x/posts/keeper"
	"github.com/warmage-sports/mage/x/posts/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/warmage-sports/mage/cosmwasm"
)

var _ cosmwasm.Querier = PostsWasmQuerier{}

type PostsWasmQuerier struct {
	postsKeeper postskeeper.Keeper
	cdc         codec.Codec
}

func NewPostsWasmQuerier(postsKeeper postskeeper.Keeper, cdc codec.Codec) PostsWasmQuerier {
	return PostsWasmQuerier{postsKeeper: postsKeeper, cdc: cdc}
}

func (querier PostsWasmQuerier) QueryCustom(ctx sdk.Context, data json.RawMessage) ([]byte, error) {
	var query types.PostsQuery
	err := json.Unmarshal(data, &query)

	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	var response []byte
	switch {
	case query.SubspacePosts != nil:
		if response, err = querier.handleSubspacePostsRequest(ctx, *query.SubspacePosts); err != nil {
			return nil, err
		}
	case query.SectionPosts != nil:
		if response, err = querier.handleSectionPostsRequest(ctx, *query.SectionPosts); err != nil {
			return nil, err
		}
	case query.Post != nil:
		if response, err = querier.handlePostRequest(ctx, *query.Post); err != nil {
			return nil, err
		}
	case query.PostAttachments != nil:
		if response, err = querier.handlePostAttachmentsRequest(ctx, *query.PostAttachments); err != nil {
			return nil, err
		}
	case query.PollAnswers != nil:
		if response, err = querier.handlePollAnswersRequest(ctx, *query.PollAnswers); err != nil {
			return nil, err
		}
	default:
		return nil, sdkerrors.ErrInvalidRequest
	}

	return response, nil
}

func (querier PostsWasmQuerier) handleSubspacePostsRequest(ctx sdk.Context, request json.RawMessage) (bz []byte, err error) {
	var PostsReq types.QuerySubspacePostsRequest
	err = querier.cdc.UnmarshalJSON(request, &PostsReq)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	PostsResponse, err := querier.postsKeeper.SubspacePosts(sdk.WrapSDKContext(ctx), &PostsReq)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	bz, err = querier.cdc.MarshalJSON(PostsResponse)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return bz, nil
}

func (querier PostsWasmQuerier) handleSectionPostsRequest(ctx sdk.Context, request json.RawMessage) (bz []byte, err error) {
	var incomingDtagReq types.QuerySectionPostsRequest
	err = querier.cdc.UnmarshalJSON(request, &incomingDtagReq)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	SectionPostsResponse, err := querier.postsKeeper.SectionPosts(sdk.WrapSDKContext(ctx), &incomingDtagReq)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	bz, err = querier.cdc.MarshalJSON(SectionPostsResponse)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return bz, nil
}

func (querier PostsWasmQuerier) handlePostRequest(ctx sdk.Context, request json.RawMessage) (bz []byte, err error) {
	var chainLinkReq types.QueryPostRequest
	err = querier.cdc.UnmarshalJSON(request, &chainLinkReq)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	PostResponse, err := querier.postsKeeper.Post(sdk.WrapSDKContext(ctx), &chainLinkReq)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	bz, err = querier.cdc.MarshalJSON(PostResponse)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

func (querier PostsWasmQuerier) handlePostAttachmentsRequest(ctx sdk.Context, request json.RawMessage) (bz []byte, err error) {
	var PostAttachmentsReq types.QueryPostAttachmentsRequest
	err = querier.cdc.UnmarshalJSON(request, &PostAttachmentsReq)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	PostAttachmentsResponse, err := querier.postsKeeper.PostAttachments(sdk.WrapSDKContext(ctx), &PostAttachmentsReq)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	bz, err = querier.cdc.MarshalJSON(PostAttachmentsResponse)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return bz, nil
}

func (querier PostsWasmQuerier) handlePollAnswersRequest(ctx sdk.Context, request json.RawMessage) (bz []byte, err error) {
	var applicationReq types.QueryPollAnswersRequest
	err = querier.cdc.UnmarshalJSON(request, &applicationReq)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	applicationLinkByChainIDResponse, err := querier.postsKeeper.PollAnswers(
		sdk.WrapSDKContext(ctx),
		&applicationReq,
	)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	bz, err = querier.cdc.MarshalJSON(applicationLinkByChainIDResponse)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return bz, nil
}
