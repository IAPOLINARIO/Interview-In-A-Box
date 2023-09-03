import {useStateValue} from "../../StateProvider";
import {Navigate, useOutlet} from "react-router-dom";
import Sidebar from "../../components/Sidebar/Sidebar";
import Footer from "../../components/Footer/Footer";
import {useOktaAuth} from "@okta/okta-react";
import {useEffect, useState} from "react";

const PrivateMainLayout = () => {

    const { oktaAuth, authState } = useOktaAuth();
    const { setUser } = useStateValue();
    const outlet = useOutlet()

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


    if (!authState || !authState.isAuthenticated) {
        return <Navigate to="/"/>
    }

    return (
        <div>
            <Sidebar/>
            <main className="p-5 sm:ml-60 min-h-[calc(100vh-5rem)] h-[calc(100vh-5rem)] overflow-auto mt-12">
                {outlet}
            </main>
            <Footer />
        </div>
    )
}

export default PrivateMainLayout