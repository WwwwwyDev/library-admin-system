import request from '@/utils/request'

export function login(data) {
  return request({
    url: '/verify/login',
    method: 'post',
    data
  })
}

export function getInfo(token) {
  return request({
    url: '/djwt',
    method: 'get',
  })
}

export function logout() {
  return request({
    url: '/verify/loginOut',
    method: 'post'
  })
}

export function refreshToken() {
  return request({
    url: '/rjwt',
    method: 'get'
  })
}

export function getUserById(id){
  return request({
    url: '/search/userId/'+String(id),
    method: 'get'
  })
}
