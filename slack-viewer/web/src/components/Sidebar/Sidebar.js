import React, {useContext, useState} from 'react';
import {Link, useLocation} from "react-router-dom";
import SidebarUserMenu from "./SidebarUserMenu";
import SidebarToggle from "./SidebarToggle";
import {navigation} from "../../infrastructure/navigation/navigation";
import DebouncedInput from "../Inputs/DebouncedInput";
import DatePicker from 'react-datepicker'
import {GlobalContext} from "../../infrastructure/context/global.context";

const Sidebar = () => {
    const [userMenuOpen, toggleUserMenu] = useState(false);
    const [sidebarOpen, toggleSidebarOpen] = useState(false);
    const {searchByUser, setSearchByUser, dateRange, setDateRange} = useContext(GlobalContext);

    const [from, to] = dateRange
    const location = useLocation()

    const doSetSearchByUser = (term) => {
        setSearchByUser(term);
    }

    const doSetDateRange = (dateRange) => {
        setDateRange(dateRange);
    }

    return (
        <div>
            <nav
                className="fixed top-0 drop-shadow z-50 w-full bg-white">
                <div className="px-3 py-2 lg:px-5 lg:pl-3">
                    <div className="flex items-center justify-between">
                        <SidebarToggle toggleSidebarOpen={toggleSidebarOpen} sidebarOpen={sidebarOpen}/>
                        <div className="flex items-center header space-x-2">
                            <DebouncedInput
                                value={searchByUser ?? ''}
                                onChange={(value) => doSetSearchByUser(String(value))}
                                placeholder="Search user by name or username"
                                label="Search users"
                                classes="min-w-[500px]"
                                debounce={500}
                            />
                            <DatePicker
                                selectsRange
                                dateFormat="MM/yyyy"
                                startDate={from}
                                endDate={to}
                                onChange={(update) => {
                                    doSetDateRange(update)
                                }}
                                className="py-1.5 px-2 text-gray-900 sm:text-sm rounded-lg  block border border-gray-400 focus:ring-emerald-600 focus:border-emerald-600 focus:outline-fuchsia-800 min-w-[225px]"
                                isClearable={true}
                                placeholderText="Select a date range"
                                showMonthYearPicker
                            />
                        </div>
                        <SidebarUserMenu toggleUserMenu={toggleUserMenu} userMenuOpen={userMenuOpen}
                                         placeholder="Search users" label="Search users" debounce={300}/>
                    </div>
                </div>
            </nav>
            <aside id="nav-sidebar"
                   className={`fixed border-r border-gray-100 top-0 left-0 z-40 w-60 h-screen pt-20 transition-transform bg-fuchsia-950 ${
                       sidebarOpen ? '-translate-x-full' : ' sm:translate-x-0'
                   }`}
                   aria-label="Sidebar">
                <div className="h-full px-3 pt-1 pb-4 overflow-y-auto bg-fuchsia-950">
                    <ul className="space-y-1.5">
                        {
                            navigation.map((navItem) => (
                                <li key={navItem.path}>
                                    <Link
                                        to={navItem.path}
                                        className={`flex items-center p-1.5 text-normal font-normal text-fuchsia-50 rounded-lg
                                        focus:outline-fuchsia-400 
                                        ${location.pathname === navItem.path ? 'bg-fuchsia-900 hover:bg-fuchsia-900 text-fuchsia-50' : 'bg-transparent hover:bg-fuchsia-800'}`}>
                                        {navItem.icon}
                                        <span className="ml-2 text-sm font-semibold">
                                            {navItem.label}
                                        </span>
                                    </Link>
                                </li>
                            ))
                        }
                    </ul>
                </div>
            </aside>
        </div>
    );
};

export default Sidebar;
