import {
  login,
  logout,
  getInfo,
  refreshToken
} from '@/api/user'
import {
  getToken,
  setToken,
  removeToken,
  setRefreshAfter,
  getRefreshAfter,
  getAccessExpire,
  setAccessExpire
} from '@/utils/auth'
import {
  resetRouter
} from '@/router'
import {
  userAvatar
} from '@/settings'
const getDefaultState = () => {
  return {
    token: getToken(),
    accessExpire: 0,
    refreshAfter: 0,
    id: 0,
    name: '',
    avatar: '',
    info: '',
    roles: [],
  }
}

const state = getDefaultState()

const getters = {
  name: state => state.name,
  avatar: state => state.avatar,
  info: state => state.info,
  id: state => state.id,
  roles: state => state.roles
}

const mutations = {
  RESET_STATE: (state) => {
    Object.assign(state, getDefaultState())
  },
  SET_TOKEN: (state, token) => {
    state.token = token
  },
  SET_ID: (state, id) => {
    state.id = id
  },
  SET_NAME: (state, name) => {
    state.name = name
  },
  SET_INFO: (state, info) => {
    state.info = info
  },
  SET_AVATAR: (state, avatar) => {
    state.avatar = avatar
  },
  SET_ACCESSEXPIRE: (state, accessExpire) => {
    state.accessExpire = accessExpire
  },
  SET_REFRESHAFTER: (state, refreshAfter) => {
    state.refreshAfter = refreshAfter
  },
  SET_ROLES: (state, roles) => {
    state.roles = roles
  }
}

const actions = {
  // user login
  login({
    commit
  }, userInfo) {
    const {
      username,
      password
    } = userInfo
    return new Promise((resolve, reject) => {
      login({
        username: username.trim(),
        password: password
      }).then(response => {
        const {
          data
        } = response
        if (!data) {
          return reject(response.msg)
        }
        commit('SET_TOKEN', data.accessToken)
        commit('SET_ACCESSEXPIRE', data.accessExpire)
        commit('SET_REFRESHAFTER', data.refreshAfter)
        setRefreshAfter(data.refreshAfter)
        setAccessExpire(data.accessExpire)
        setToken(data.accessToken)
        resolve()
      }).catch(error => {
        reject(error)
      })
    })
  },

  // get user info
  getInfo({
    commit,
    state
  }) {
    return new Promise((resolve, reject) => {
      getInfo().then(response => {
        const {
          data
        } = response
        if (!data) {
          return reject('验证失败,请重新登录')
        }
        const {
          id,
          username,
          avatar,
          info,
          roles
        } = data
        commit('SET_ROLES', roles)
        commit('SET_ID', id)
        commit('SET_NAME', username)
        commit('SET_AVATAR', avatar === "" ? userAvatar : avatar)
        commit('SET_INFO', info)
        resolve(data)
      }).catch(error => {
        reject(error)
      })
    })
  },

  // user logout
  logout({
    commit,
    state
  }) {
    return new Promise((resolve, reject) => {
      logout(state.token).then(() => {
        removeToken() // must remove  token  first
        resetRouter()
        commit('RESET_STATE')
        resolve()
      }).catch(error => {
        reject(error)
      })
    })
  },

  // remove token
  resetToken({
    commit
  }) {
    return new Promise(resolve => {
      removeToken() // must remove  token  first
      commit('RESET_STATE')
      resolve()
    })
  },
  refreshToken({
    commit,
    state
  }) {
    return new Promise((resolve, reject) => {
      let timeNow = Math.round(new Date().getTime() / 1000)
      let accessExpire = getAccessExpire()
      let refreshAfter = getRefreshAfter()
      if (timeNow > accessExpire) {
        return reject('会话失效,请重新登录')
      }
      if (timeNow > refreshAfter) {
        refreshToken().then(response => {
          const {
            accessExpire,
            accessToken,
            refreshAfter
          } = response.data
          commit('SET_TOKEN', accessToken)
          commit('SET_ACCESSEXPIRE', accessExpire)
          commit('SET_REFRESHAFTER', refreshAfter)
          setRefreshAfter(refreshAfter)
          setAccessExpire(accessExpire)
          setToken(accessToken)
        }).catch(error => {
          reject(error)
        })
      }
      resolve()
    })
  }
}

export default {
  namespaced: true,
  state,
  getters,
  mutations,
  actions
}
