import request from '@/utils/request'

export function getTypes(params) {
  return request({
    url: '/book/type',
    method: 'get',
    params
  })
}
