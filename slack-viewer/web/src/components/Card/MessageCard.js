import React, {useContext} from 'react';
import PropTypes from "prop-types";
import {GlobalContext} from "../../infrastructure/context/global.context";
import {highlightPortionOfTheText} from "../../infrastructure/utils/utils";
import {Person} from "@material-ui/icons";

const MessageAvatar = ({avatar}) => {
    if (!avatar) {
        return (
            <div className="relative w-12 h-12 overflow-hidden bg-gray-100 rounded-full text-xl flex items-center justify-center">
                <Person className="text-gray-400" fontSize="large"/>
            </div>
        )
    }
    return (
        <img className=" w-12 h-12  p-1 rounded-full ring-2 ring-fuchsia-700"
             src={avatar} alt="Bordered avatar"/>
    )
}


const MessageCard = ({message, user, userId, humanReadableDate}) => {
    const {searchByUser} = useContext(GlobalContext);
    return (
        <div className="px-3 py-2.5 bg-white drop-shadow flex flex-col space-y-3">
            <div className="flex flex-row items-center space-x-2">
                <MessageAvatar avatar={user.avatar}/>
                <div className="flex flex-col space-y-0.5">
                    <p className="text-xs font-bold">{user?.fullName ? <span
                        dangerouslySetInnerHTML={{__html: highlightPortionOfTheText(user.fullName, searchByUser)}} /> : '-'}</p>
                    <p className="text-xs font-semibold">{userId ?? '-'}</p>
                    <p className="text-xs font-normal italic">{humanReadableDate ?? '-'}</p>
                </div>
            </div>
            <div className="text-sm w-full">
                <figure>
                    <pre role="textbox" style={{width: '160ch', whiteSpace: 'pre-wrap'}}>
                        {message}
                    </pre>
                </figure>

            </div>
        </div>
    );
};

MessageCard.propTypes = {
    message: PropTypes.string,
    timestamp: PropTypes.number,
    userId: PropTypes.string,
    humanReadableDate: PropTypes.string,
    user: PropTypes.object
}

export default MessageCard;
