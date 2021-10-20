import request from '@/utils/request'

export function getLends(params) {
  return request({
    url: '/lend',
    method: 'get',
    params
  })
}


export function deleteLend(id) {
  return request({
    url: '/lend/'+String(id),
    method:'delete'
  })
}


export function addLend(data) {
  return request({
    url: '/lend',
    method: 'post',
    data
  })
}
