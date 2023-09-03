import React, {useContext, useEffect, useState} from 'react';
import EmptyStateHandler from "../../components/EmptyStateHandler/EmptyStateHandler";
import {useGroupDirectMessages} from "./hooks/useGroupDirectMessages";
import {GlobalContext} from "../../infrastructure/context/global.context";
import MessageCard from "../../components/Card/MessageCard";
import TitleWithCountHeader from "../../components/TitleWithCountHeader";
import TabItem from "../../components/Tabs/TabItem";
import Tabs from "../../components/Tabs/Tabs";

const GroupDirectMessagesContent = () => {
    const {searchByUser, dateRange: [from, to]} = useContext(GlobalContext);
    const {groupDirectMessages, groups, errorMessage} = useGroupDirectMessages({searchByUser, from, to});
    const [currentTab, setCurrentTab] = useState(null);

    useEffect(() => {
        if (groupDirectMessages.length > 0 && currentTab === null) {
            setCurrentTab(groupDirectMessages[0].groupName);
        }
        return () => {
            // setCurrentTab(null);
        }
    }, [groupDirectMessages])


    return (
        <div>
            <TitleWithCountHeader title="Groups" count={groupDirectMessages.length} />
            <EmptyStateHandler data={groupDirectMessages} errorMessage={errorMessage} keyword={'group direct messages'}>
                <div className="flex flex-col overflow-auto z-10 space-y-2">
                    <Tabs>
                        {groups?.map((group, index) => (
                            <TabItem key={index} text={group} onClick={() => setCurrentTab(group)} active={currentTab === group}/>
                        ))}
                    </Tabs>
                    {groupDirectMessages.filter((data, index) => data.groupName === currentTab).map((gdm, index) => (
                        <div key={index}>
                            <MessageCard message={gdm.message}
                                         timestamp={gdm.timestamp} userId={gdm.userId}
                                         humanReadableDate={gdm.humanReadableDate}
                                         user={gdm.user}/>
                        </div>
                    ))}
                </div>
            </EmptyStateHandler>
        </div>
    );
};

export default GroupDirectMessagesContent;
