import request from '@/utils/request'

export function getVips(params) {
  return request({
    url: '/vip',
    method: 'get',
    params
  })
}

export function deleteVip(id) {
  return request({
    url: '/vip/'+String(id),
    method:'delete'
  })
}


export function addVip(data) {
  return request({
    url: '/vip',
    method: 'post',
    data
  })
}

export function editVip(id, data) {
  return request({
    url: '/vip/'+String(id),
    method: 'put',
    data
  })
}
