import * as qiniu from 'qiniu-js'
import {
  qiniuDomain
} from '@/settings'
export const domain = qiniuDomain

export function upload(file, token) {
  let key = null
  let config = {
    useCdnDomain: true,
    region: qiniu.region.z0,
    uptoken: token,
    domain: domain, //配置好的七牛云域名  如   https://cdn.qniyun.com/
    unique_names: true,
  };
  let putExtra = {
    // fname: file.name,
    // params: {},
    // mimeType: [] || null
    ...config,
  }
  return qiniu.upload(file, key, token, putExtra, config)
}

export function uploadAsync(file, token) {
  return new Promise((resolve, reject) => {
    const observable = upload(file, token)
    const subscription = observable.subscribe({
      next(res) {
        //进度
      },
      error(err) {
        reject(err)
      },
      complete(res) {
        resolve(domain + res.hash)
      }
    })
  })
}


