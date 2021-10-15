import request from '@/utils/request'

export function getLoginStatus(){
  return request({
    url:'/search/loginStatus',
    method:'get'
  })
}
export function getUsers(params) {
  return request({
    url: '/user',
    method: 'get',
    params
  })
}

export function deleteUser(id) {
  return request({
    url: '/user/'+String(id),
    method:'delete'
  })
}


export function addUser(data) {
  return request({
    url: '/user',
    method: 'post',
    data
  })
}

export function editUser(id, data) {
  return request({
    url: '/user/'+String(id),
    method: 'put',
    data
  })
}

export function getRoles(){
  return request({
    url: '/search/roles',
    method: 'get'
  })
}