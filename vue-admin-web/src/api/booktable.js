import request from '@/utils/request'

export function getBooks(params) {
  return request({
    url: '/book',
    method: 'get',
    params
  })
}

export function editBook(id,data){
  return request({
    url: '/book/'+String(id),
    method: 'put',
    data
  })
}

export function deleteBook(id) {
  return request({
    url: '/book/'+String(id),
    method:'delete'
  })
}


export function addBook(data) {
  return request({
    url: '/book',
    method: 'post',
    data
  })
}
