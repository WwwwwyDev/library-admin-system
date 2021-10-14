import request from '@/utils/request'

export function getQiniuToken() {
  return request({
    url: '/qiniuToken',
    method: 'get',
  })
}