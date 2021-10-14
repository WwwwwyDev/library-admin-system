import * as qiniu from 'qiniu-js'

export const domain = "http://file.wwywwy.top/"

export function upload(file,token) {
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
