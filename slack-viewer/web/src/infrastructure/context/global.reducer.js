import {GlobalSelectors} from "./global.selectors";
const globalSelectorReducer = (state, action) => {
    switch (action.type) {
        case GlobalSelectors.SET_SEARCH_BY_USER:
            return {
                ...state,
                searchByUser: action.payload ?? undefined
            }
        case GlobalSelectors.SET_DATE_RANGE:
            return {
                ...state,
                dateRange: action.payload ?? []
            }
        default:
            return {
                ...state
            }
    }
}

export { globalSelectorReducer }
