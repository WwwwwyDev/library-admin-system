import request from '@/utils/request'

export function getBooks(params) {
  return request({
    url: '/book',
    method: 'get',
    params
  })
}