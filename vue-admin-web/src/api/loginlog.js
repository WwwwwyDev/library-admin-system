import request from '@/utils/request'

export function getLoginLogs(params) {
  return request({
    url: '/systemlog',
    method: 'get',
    params
  })
}

export function deleteLoginLogs() {
  return request({
    url: '/systemlog',
    method:'delete'
  })
}
