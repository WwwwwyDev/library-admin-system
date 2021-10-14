import request from '@/utils/request'

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
