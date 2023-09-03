import {useOktaAuth} from "@okta/okta-react";
import {Navigate} from "react-router-dom";
import {useStateValue} from "../StateProvider";
import {useEffect} from "react";

const SignIn = () => {
    const { oktaAuth, authState } = useOktaAuth();
    const { setUser } = useStateValue();
    
    const onSuccess = async (tokens) => {
        await oktaAuth.handleLoginRedirect(tokens);
    }
    const onError = (err) => {
        console.log('Sign in error:', err);
    };

    useEffect(() => {
        if (!authState || !authState.isAuthenticated) {
            // When user isn't authenticated, forget any user info
            setUser(null);
        } else {
            oktaAuth.getUser().then((info) => {
                setUser(info);
            }).catch((err) => {
                console.error(err);
            });
        }
    }, [authState, oktaAuth]);

    if(!authState) return <div>Loading...</div>;

    return <Navigate to={{ pathname: '/admin' }}/>
};

export default SignIn;
