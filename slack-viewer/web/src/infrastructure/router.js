import PrivateMainLayout from "./layout/PrivateMainLayout";
import Dashboard from "../pages/Dashboard";
import PublicChannels from "../pages/PublicChannels";
import PrivateChannels from "../pages/PrivateChannels";
import DirectMessages from "../pages/DirectMessages";
import GroupDirectMessages from "../pages/GroupDirectMessages";
import SignIn from "../pages/SignIn";
import PublicMainLayout from "./layout/PublicMainLayout";
import Home from "../pages/Home";
import {LoginCallback} from "@okta/okta-react";

export const router = [
    {
        path: '/',
        element: <PublicMainLayout />,
        children: [
            {
                index: true,
                element: <Home />
            }
        ]
    },
    {
        path: '/admin',
        element: <PrivateMainLayout />,
        children: [
            {
                index: true,
                element: <Dashboard />
            },
            {
                path: 'public-channels',
                element: <PublicChannels />
            },
            {
                path: 'private-channels',
                element: <PrivateChannels />
            },
            {
                path: 'direct-messages',
                element: <DirectMessages />
            },
            {
                path: 'group-direct-messages',
                element: <GroupDirectMessages />
            }
        ]
    },
    {
        path: 'auth/sign-in',
        element: <SignIn />
    },
    {
        path: 'auth/sign-in/callback',
        element: <LoginCallback />
    }
]
