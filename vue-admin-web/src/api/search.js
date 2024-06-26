import request from '@/utils/request'

export function getTypesByNameLike(name) {
  return request({
    url: '/search/book/type/'+String(name),
    method: 'get'
  })
}

export function getAllTypes() {
  return request({
    url: '/search/book/type',
    method: 'get'
  })
}

export function getTypeById(id) {
  return request({
    url: '/search/book/typeId/'+String(id),
    method: 'get'
  })
}


export function getVipByCardNumberLike(cardNumber) {
  return request({
    url:"/search/vip/"+String(cardNumber),
    method: 'get'
  })
}

export function getBookByNameLike(name){
  return request({
    url:"/search/book/"+String(name),
    method: 'get'
  })
}