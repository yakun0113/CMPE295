const state = {	
	walmartList: [
	
	],
    targetList: [
		
	],
    walgreensList: [
		
	],
	storeLocations:{
		"walmart":{},
		'target':{},
		"walgreens":{}
	},
}

const mutations = {

	'SET_PRODUCTS' (state, products) {
		state.walmartList = products[0].products;
        state.targetList = products[1].products;
        state.walgreensList = products[2].products;
	
		state.storeLocations['walmart'] = products[0].location;
		state.storeLocations['target'] = products[1].location;
		state.storeLocations['walgreens'] = products[2].location;
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
    },
	walmart_sort_by_price: () => {
		let sortedList = state.walmartList
		sortedList.sort((a, b) => (a.price.substring(1)-'0' > b.price.substring(1)-'0') ? 1 : -1)
		return sortedList
	},
	target_sort_by_price: () => {
		let sortedList = state.targetList
		sortedList.sort((a, b) => (a.price.substring(1)-'0' > b.price.substring(1)-'0') ? 1 : -1)
		return sortedList
	},
	walgreens_sort_by_price: () => {
		let sortedList = state.walgreensList
		sortedList.sort((a, b) => (a.price.substring(1)-'0' > b.price.substring(1)-'0') ? 1 : -1)
		return sortedList
	},
	walmart_sort_by_rating: () => {
		let sortedList = state.walmartList
		sortedList.sort((a, b) => (a.rating.substring(0,3)-'0' < b.rating.substring(0,3)-'0') ? 1 : -1)
		return sortedList
	},
	target_sort_by_rating: () => {
		let sortedList = state.targetList
		sortedList.sort((a, b) => (a.rating.substring(0,3)-'0' < b.rating.substring(0,3)-'0') ? 1 : -1)
		return sortedList
	},
	walgreens_sort_by_rating: () => {
		let sortedList = state.walgreensList
		sortedList.sort((a, b) => (a.rating.substring(0,3)-'0' < b.rating.substring(0,3)-'0') ? 1 : -1)
		return sortedList
	},
	store_locations: (state) => {
		return state.storeLocations;
	},
}

export default {
	state,
	mutations,
	actions,
	getters
}