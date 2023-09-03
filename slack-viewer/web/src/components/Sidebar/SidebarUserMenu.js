import React from 'react';
import PropTypes from 'prop-types';
import { PersonOutline } from "@material-ui/icons";
import PrimaryContainedButton from "../Buttons/PrimaryContainedButton";
import {useOktaAuth} from "@okta/okta-react";
import {useStateValue} from "../../StateProvider";

const SidebarUserMenu = ({ toggleUserMenu, userMenuOpen }) => {
    const { oktaAuth, authState } = useOktaAuth();
    const { user } = useStateValue();

    const handleLogout = async () => oktaAuth.signOut();

    return (
        <div className="flex items-center">
            <div className="flex items-center ml-3 relative">
                <div>
                    <button
                        type="button"
                        className="flex text-sm bg-fuchsia-50 rounded-full ring-2 ring-fuchsia-950/75"
                        aria-expanded="false"
                        onClick={() => {
                            toggleUserMenu(!userMenuOpen)
                        }}
                    >
                        <span className="sr-only">Open user menu</span>
                        <PersonOutline />
                    </button>
                </div>
                <div
                    className={`z-50 my-4 text-base list-none bg-white divide-y divide-gray-100 rounded shadow ${
                        userMenuOpen
                            ? 'visible absolute block right-0 top-5'
                            : 'hidden'
                    }`}
                >
                    <div className="px-4 py-3" role="none">
                        <p
                            className="text-sm text-gray-900"
                            role="none"
                        >
                            {user?.name}
                        </p>
                        <p
                            className="text-sm font-medium text-fuchsia-950 truncate"
                            role="none"
                        >
                            {user?.email}
                        </p>
                    </div>
                    <ul className="py-1" role="none">
                        {/*<li>*/}
                        {/*    <Link*/}
                        {/*        to="/dashboard"*/}
                        {/*        className="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"*/}
                        {/*        role="menuitem"*/}
                        {/*    >*/}
                        {/*        Profile*/}
                        {/*    </Link>*/}
                        {/*</li>*/}
                        {/*<li>*/}
                        {/*    <Link*/}
                        {/*        to="/dashboard"*/}
                        {/*        className="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"*/}
                        {/*        role="menuitem"*/}
                        {/*    >*/}
                        {/*        Settings*/}
                        {/*    </Link>*/}
                        {/*</li>*/}
                        <li className="px-4 py-2">
                            <PrimaryContainedButton size="xs" fullWidth onClick={() => handleLogout()}>
                                Sign out
                            </PrimaryContainedButton>
                        </li>
                    </ul>
                </div>
            </div>
        </div>
    );
};

SidebarUserMenu.propTypes = {
    toggleUserMenu: PropTypes.func.isRequired,
    userMenuOpen: PropTypes.bool.isRequired,
};

export default SidebarUserMenu;
