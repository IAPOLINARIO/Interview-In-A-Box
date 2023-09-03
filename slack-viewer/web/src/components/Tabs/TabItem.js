import React from 'react';
import PropTypes from "prop-types";
import {truncate} from "../../infrastructure/utils/utils";

const TabItem = ({text, active = false, ...props}) => {
    return (
        <li
            className={`inline-block p-4 border-b-2 rounded-t-lg hover:text-fuchsia-700 hover:border-fuchsia-700 
            cursor-pointer ${active ? ' active border-fuchsia-700 text-fuchsia-700' : ' border-transparent '}`}
            {...props}
        >
            {truncate(text)}
        </li>
    );
};

TabItem.propTypes = {
    text: PropTypes.string.isRequired,
    active: PropTypes.bool
}

export default TabItem;
