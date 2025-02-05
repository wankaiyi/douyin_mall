package utils

import (
	"context"
	"fmt"
	"github.com/bytedance/gopkg/cloud/metainfo"
)

func EnsurePersistentValue(ctx context.Context, key string, isString bool) context.Context {
	value := ctx.Value(key)

	if value != nil {
		if isString {
			ctx = metainfo.WithPersistentValue(ctx, key, value.(string))
		} else {
			ctx = metainfo.WithPersistentValue(ctx, key, fmt.Sprintf("%v", value))
		}
	} else {
		persistentValue, ok := metainfo.GetPersistentValue(ctx, key)
		if ok {
			ctx = metainfo.WithPersistentValue(ctx, key, persistentValue)
		}
	}

	return ctx
}
