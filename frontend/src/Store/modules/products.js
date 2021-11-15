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

	storeLinks:{
		"walmart":"",
		'target':"",
		"walgreens":"",
	},
}

const mutations = {

	'SET_PRODUCTS' (state, products) {
		state.walmartList = products.find(function(o){ return o.store=="Walmart"}).products;
        state.targetList = products.find(function(o){ return o.store=="Target"}).products;
        state.walgreensList = products.find(function(o){ return o.store=="Walgreens"}).products;
	
		state.storeLocations['walmart'] = products.find(function(o){ return o.store=="Walmart"}).location;
		state.storeLocations['target'] = products.find(function(o){ return o.store=="Target"}).location;
		state.storeLocations['walgreens'] = products.find(function(o){ return o.store=="Walgreens"}).location;

		state.storeLinks['walmart'] = products.find(function(o){ return o.store=="Walmart"}).website;
		state.storeLinks['target'] = products.find(function(o){ return o.store=="Target"}).website;
		state.storeLinks['walgreens'] = products.find(function(o){ return o.store=="Walgreens"}).website;
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

	store_locations: (state) => {
		return state.storeLocations;
	},

	store_links: (state) =>{
		return state.storeLinks;
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
}

export default {
	state,
	mutations,
	actions,
	getters
}