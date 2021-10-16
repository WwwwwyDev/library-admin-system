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

//通过jwt修改用户信息
export function editUser(data){
  return request({
    url: 'verify/editUser',
    method: 'post',
    data
  })
}
//通过jwt修改用户密码
export function changePassword(data){
  return request({
    url: 'verify/changePassword',
    method: 'post',
    data
  })
}