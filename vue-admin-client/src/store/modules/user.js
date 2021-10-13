import {
  login,
  logout,
  getInfo
} from '@/api/user'
import {
  getToken,
  setToken,
  removeToken,
} from '@/utils/auth'
import {
  resetRouter
} from '@/router'

const getDefaultState = () => {
  return {
    token: getToken(),
    accessExpire: 0,
    refreshAfter: 0,
    name: '',
    avatar: '',
    info: '',
  }
}

const state = getDefaultState()

const getters = {
  name: state => state.name,
  avatar: state => state.avatar,
  info: state=> state.info,
}

const mutations = {
  RESET_STATE: (state) => {
    Object.assign(state, getDefaultState())
  },
  SET_TOKEN: (state, token) => {
    state.token = token
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
  SET_REFRESHAFTER: (state, info) => {
    state.info = info
  },
  SET_ISLOAD: (state, isLoad) => {
    state.isLoad = isLoad
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
          username,
          avatar,
          info
        } = data
        commit('SET_NAME', username)
        commit('SET_AVATAR', avatar)
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
    commit
  }) {
    return new Promise(resolve => {
      let timeNow =	Math.round(new Date().getTime()/1000)
      console.log(timeNow)
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
