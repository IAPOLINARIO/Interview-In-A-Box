import {globalSelectorReducer} from "./global.reducer";
import React, {useReducer, createContext} from "react";
import {GlobalSelectors} from "./global.selectors";
import PropTypes from "prop-types";

const initialState = {
    searchByUser: undefined,
    dateRange: []
}

export const GlobalContext = createContext(initialState);

const init = (initialState) => {
    return {
        searchByUser: initialState.searchByUser,
        dateRange: initialState.dateRange
    }
}

export const GlobalContextProvider = ({ children }) => {
    const [state, dispatch] = useReducer(globalSelectorReducer, initialState, init)

    const setSearchByUser = (searchTerm) => {
        dispatch({
            type: GlobalSelectors.SET_SEARCH_BY_USER,
            payload: searchTerm
        })
    }

    const setDateRange = (dateRange) => {
        dispatch({
            type: GlobalSelectors.SET_DATE_RANGE,
            payload: dateRange
        })
    }

    return (
        <GlobalContext.Provider value={{
            searchByUser: state.searchByUser,
            dateRange: state.dateRange,
            setSearchByUser,
            setDateRange
        }}>
            {children}
        </GlobalContext.Provider>
    )
}

GlobalContextProvider.propTypes = {
    children: PropTypes.node.isRequired
}
