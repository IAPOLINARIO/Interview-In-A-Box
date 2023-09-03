import React, {createContext, useContext, useReducer} from "react";
import reducer, {actionTypes, initialState} from "./reducer";

export const StateContext = createContext();

const init = (initialState) => {
    return {user: initialState.user}
}

export const StateProvider = ({ children }) => {
    const [state, dispatch] = useReducer(reducer, initialState, init);
    const setUser = (userData) => {
        dispatch({
            type: actionTypes.SET_USER,
            payload: userData
        })
    }

    return (
        <StateContext.Provider value={{
            user: state.user,
            setUser
        }}>
            {children}
        </StateContext.Provider>
    )
};

export const useStateValue = () => useContext(StateContext);