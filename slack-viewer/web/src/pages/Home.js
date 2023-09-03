import React from 'react';
import PrimaryContainedButton from "../components/Buttons/PrimaryContainedButton";
import {useOktaAuth} from "@okta/okta-react";

const Home = () => {
    const { oktaAuth, authState } = useOktaAuth();
    const login = async () => oktaAuth.signInWithRedirect({ originalUri: '/auth/sign-in'});
    return (
        <div className="relative">
            <div className="h-3 w-full absolute top-0 left-0 right-0 z-20  bg-fuchsia-900"></div>
            <div  className="flex flex-col space-y-5 items-center justify-center max-w-4xl mx-auto">
                <div className="-translate-y-12 flex flex-col items-center justify-center h-screen">
                    <h4 className="text-5xl font-bold mb-1">Welcome</h4>
                    <div className="flex flex-col text-2xl mb-10">
                        <span className="self-center font-semibold whitespace-nowrap text-gray-900">
                        Slack<span className="text-fuchsia-800">Viewer</span>
                        </span>
                    </div>
                    <PrimaryContainedButton size="lg" fullWidth onClick={() => login()}>Sign In</PrimaryContainedButton>
                </div>
            </div>
        </div>
    );
};

export default Home;
