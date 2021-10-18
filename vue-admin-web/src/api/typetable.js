import request from '@/utils/request'

export function getTypes(params) {
  return request({
    url: '/book/type',
    method: 'get',
    params
  })
}


export function editType(id,data){
  return request({
    url: '/book/type/'+String(id),
    method: 'put',
    data
  })
}

export function deleteType(id) {
  return request({
    url: '/book/type/'+String(id),
    method:'delete'
  })
}


export function addType(data) {
  return request({
    url: '/book/type',
    method: 'post',
    data
  })
}
