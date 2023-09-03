import React from 'react';
import PropTypes from "prop-types";

const Tabs = ({ children }) => {
    return (
        <ul className="flex text-sm font-medium text-center text-gray-500 border-b border-gray-200 space-x-2 overflow-x-auto">
            {children}
        </ul>
    );
};

Tabs.propTypes = {
    children: PropTypes.node,
}

export default Tabs;

