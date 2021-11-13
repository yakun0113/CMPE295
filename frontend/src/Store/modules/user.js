const state = {	
    isLoggedin: false,
    currentUser: null,  
    watchlist:[],
    inWatchlist:false,
}

const mutations = {

	'SIGN_IN' (state, user) {
		state.isLoggedin = true;
        state.currentUser = user;

	},

    'SIGN_OUT'(state){
        state.isLoggedin = false;
        state.currentUser = "";
    },

    'SET_WATCHLIST'(state, result){
        state.watchlist = result;
        state.inWatchlist = true;
    },
}

const actions = {
    signIn({commit},user){
        commit('SIGN_IN',user)
    },

    signOut({commit}){
        commit('SIGN_OUT')
    },

    setWatchlist({commit}, result){
        commit('SET_WATCHLIST', result)
    },
}

const getters = {
	user: (state) => {
		return state.currentUser;
    },
    logged:(state) => {
        return state.isLoggedin;
    },

    watchlist:(state) => {
        return state.watchlist;
    },

    inWatchlist:(state) => {
        return state.inWatchlist;
    }
}

export default {
	state,
	mutations,
	actions,
	getters
}