import Cookies from 'js-cookie'
const Storage = window.localStorage
const TokenKey = 'authorization'
const AccessExpireKey = 'accessExpire'
const RefreshAfterKey = 'refreshAfter'

export function getToken() {
  return Storage.getItem(TokenKey)
}

export function setToken(token) {
  return Storage.setItem(TokenKey, token)
}

export function removeToken() {
  return Storage.removeItem(TokenKey)
}

export function getAccessExpire() {
  return Storage.getItem(AccessExpireKey)
}

export function setAccessExpire(accessExpire) {
  return Storage.setItem(AccessExpireKey, accessExpire)
}

export function removeAccessExpire() {
  return Storage.removeItem(AccessExpireKey)
}

export function getRefreshAfter() {
  return Storage.getItem(RefreshAfterKey)
}

export function setRefreshAfter(refreshAfter) {
  return Storage.setItem(RefreshAfterKey, refreshAfter)
}

export function removeRefreshAfter() {
  return Storage.removeItem(RefreshAfterKey)
}
