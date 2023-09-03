import React from "react";
import "./App.css";
import {useLocation, useNavigate, useRoutes} from "react-router-dom";
import {router} from "./infrastructure/router";
import {GlobalContextProvider} from "./infrastructure/context/global.context";
import {OktaAuth, toRelativeUrl} from "@okta/okta-auth-js";
import {Security} from "@okta/okta-react";

const oktaAuth = new OktaAuth({
    issuer: `https://${process.env.REACT_APP_OKTA_ORG}`,
    clientId: process.env.REACT_APP_OKTA_CLIENT_ID,
    redirectUri: window.location.origin + '/auth/sign-in/callback',
    scopes: process.env.REACT_APP_OKTA_SCOPES.split(/\s+/),
    pkce: true
});

function App() {

    let navigate = useNavigate();
    const location = useLocation();
    const element = useRoutes(router);

    const restoreOriginalUri = async (_oktaAuth, originalUri) => {
        navigate(toRelativeUrl(originalUri || "/", location.origin));
      };

    return (
         <Security oktaAuth={oktaAuth} restoreOriginalUri={restoreOriginalUri}>
            <GlobalContextProvider>
                <div className="App">
                    {element}
                </div>
            </GlobalContextProvider>
         </Security>
    );
}

export default App;