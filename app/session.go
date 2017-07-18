package app

import (
	l4g "github.com/alecthomas/log4go"
	"github.com/primefour/xserver/model"
	"github.com/primefour/xserver/utils"
	"net/http"
)

var sessionCache *utils.Cache = utils.NewLru(model.SESSION_CACHE_SIZE)

func CreateSession(session *model.Session) (*model.Session, *model.AppError) {
	if result := <-Srv.Store.Session().Save(session); result.Err != nil {
		return nil, result.Err
	} else {
		session := result.Data.(*model.Session)

		AddSessionToCache(session)

		return session, nil
	}
}

func GetSession(token string) (*model.Session, *model.AppError) {

	var session *model.Session
	if ts, ok := sessionCache.Get(token); ok {
		session = ts.(*model.Session)
	} else {
	}

	if session == nil {
		if sessionResult := <-Srv.Store.Session().Get(token); sessionResult.Err != nil {
			return nil, model.NewLocAppError("GetSession", "api.context.invalid_token.error", map[string]interface{}{"Token": token, "Error": sessionResult.Err.DetailedError}, "")
		} else {
			session = sessionResult.Data.(*model.Session)

			if session == nil || session.IsExpired() || session.Token != token {
				return nil, model.NewLocAppError("GetSession", "api.context.invalid_token.error", map[string]interface{}{"Token": token, "Error": ""}, "")
			} else {
				AddSessionToCache(session)
				return session, nil
			}
		}
	}

	if session == nil || session.IsExpired() {
		return nil, model.NewLocAppError("GetSession", "api.context.invalid_token.error", map[string]interface{}{"Token": token}, "")
	}

	return session, nil
}

func GetSessions(userId string) ([]*model.Session, *model.AppError) {
	if result := <-Srv.Store.Session().GetSessions(userId); result.Err != nil {
		return nil, result.Err
	} else {
		return result.Data.([]*model.Session), nil
	}
}

func RevokeAllSessions(userId string) *model.AppError {
	if result := <-Srv.Store.Session().GetSessions(userId); result.Err != nil {
		return result.Err
	} else {
		sessions := result.Data.([]*model.Session)

		for _, session := range sessions {
			if session.IsOAuth {

			} else {
				if result := <-Srv.Store.Session().Remove(session.Id); result.Err != nil {
					return result.Err
				}
			}

		}
	}

	ClearSessionCacheForUser(userId)

	return nil
}

func ClearSessionCacheForUser(userId string) {

	ClearSessionCacheForUserSkipClusterSend(userId)
}

func ClearSessionCacheForUserSkipClusterSend(userId string) {
	keys := sessionCache.Keys()

	for _, key := range keys {
		if ts, ok := sessionCache.Get(key); ok {
			session := ts.(*model.Session)
			if session.UserId == userId {
				sessionCache.Remove(key)
			}
		}
	}

}

func AddSessionToCache(session *model.Session) {
	sessionCache.AddWithExpiresInSecs(session.Token, session, int64(*utils.Cfg.ServiceSettings.SessionCacheInMinutes*60))
}

func SessionCacheLength() int {
	return sessionCache.Len()
}

func RevokeSessionsForDeviceId(userId string, deviceId string, currentSessionId string) *model.AppError {
	if result := <-Srv.Store.Session().GetSessions(userId); result.Err != nil {
		return result.Err
	} else {
		sessions := result.Data.([]*model.Session)
		for _, session := range sessions {
			if session.DeviceId == deviceId && session.Id != currentSessionId {
				l4g.Debug(utils.T("api.user.login.revoking.app_error"), session.Id, userId)
				if err := RevokeSession(session); err != nil {
					// Soft error so we still remove the other sessions
					l4g.Error(err.Error())
				}
			}
		}
	}

	return nil
}

func RevokeSessionById(sessionId string) *model.AppError {
	if result := <-Srv.Store.Session().Get(sessionId); result.Err != nil {
		result.Err.StatusCode = http.StatusBadRequest
		return result.Err
	} else {
		return RevokeSession(result.Data.(*model.Session))
	}
}

func RevokeSession(session *model.Session) *model.AppError {
	if session.IsOAuth {
	} else {
		if result := <-Srv.Store.Session().Remove(session.Id); result.Err != nil {
			return result.Err
		}
	}

	ClearSessionCacheForUser(session.UserId)

	return nil
}

func AttachDeviceId(sessionId string, deviceId string, expiresAt int64) *model.AppError {
	if result := <-Srv.Store.Session().UpdateDeviceId(sessionId, deviceId, expiresAt); result.Err != nil {
		return result.Err
	}

	return nil
}

func UpdateLastActivityAtIfNeeded(session model.Session) {
	now := model.GetMillis()
	if now-session.LastActivityAt < model.SESSION_ACTIVITY_TIMEOUT {
		return
	}

	if result := <-Srv.Store.Session().UpdateLastActivityAt(session.Id, now); result.Err != nil {
		l4g.Error(utils.T("api.status.last_activity.error", session.UserId, session.Id))
	}

	session.LastActivityAt = now
	AddSessionToCache(&session)
}