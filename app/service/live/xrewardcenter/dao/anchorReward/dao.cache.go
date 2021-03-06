// Code generated by $GOPATH/src/go-common/app/tool/cache/gen. DO NOT EDIT.

/*
  Package anchorReward is a generated cache proxy package.
  It is generated from:
  type _cache interface {
		// cache: -sync=true -nullcache=&model.AnchorRewardConf{ID:-1} -check_null_code=$.ID==-1
		RewardConf(c context.Context, id int64) (*model.AnchorRewardConf, error)
	}
*/

package anchorReward

import (
	"context"

	model "go-common/app/service/live/xrewardcenter/model/anchorTask"
	"go-common/library/stat/prom"
)

var _ _cache

// RewardConf get data from cache if miss will call source method, then add to cache.
func (d *Dao) RewardConf(c context.Context, id int64) (res *model.AnchorRewardConf, err error) {
	addCache := true
	res, err = d.CacheRewardConf(c, id)
	if err != nil {
		addCache = false
		err = nil
	}
	defer func() {
		if res != nil && res.ID == -1 {
			res = nil
		}
	}()
	if res != nil {
		prom.CacheHit.Incr("RewardConf")
		return
	}
	prom.CacheMiss.Incr("RewardConf")
	res, err = d.RawRewardConf(c, id)
	if err != nil {
		return
	}
	miss := res
	if miss == nil {
		miss = &model.AnchorRewardConf{ID: -1}
	}
	if !addCache {
		return
	}
	d.AddCacheRewardConf(c, id, miss)
	return
}
