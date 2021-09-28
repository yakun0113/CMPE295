const state = {	
	walmartList: [],
    targetList: [],
    walgreensList: [],
}

const mutations = {

	'SET_PRODUCTS' (state, products) {
		state.walmartList = products[0].products;
        state.targetList = products[1].products;
        state.walgreensList = products[2].products;
	},
}

const actions = {
    setProduct({commit},products){
        commit('SET_PRODUCTS',products)
    }
}

const getters = {
	walmart: (state) => {
		return state.walmartList;
	},
	target: (state) => {
		return state.targetList;
	},
    walgreens: (state) => {
        return state.walgreensList;
    }
}

export default {
	state,
	mutations,
	actions,
	getters
}