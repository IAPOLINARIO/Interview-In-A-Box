import React, {useContext, useEffect, useState} from 'react';
import {useGetDMs} from "./hooks/useDMs";
import EmptyStateHandler from "../../components/EmptyStateHandler/EmptyStateHandler";
import {GlobalContext} from "../../infrastructure/context/global.context";
import MessageCard from "../../components/Card/MessageCard";
import TitleWithCountHeader from "../../components/TitleWithCountHeader";
import TabItem from "../../components/Tabs/TabItem";
import Tabs from "../../components/Tabs/Tabs";

const DirectMessagesContent = () => {
    const { searchByUser, dateRange: [from ,to] } = useContext(GlobalContext);
    const {dms, groups, errorMessage} = useGetDMs({searchByUser, from, to});
    const [currentTab, setCurrentTab] = useState(null);


    useEffect(() => {
        if (dms.length > 0 && currentTab === null) {
            setCurrentTab(dms[0].groupName);
        }
        return () => {
            // setCurrentTab(null);
        }
    }, [dms])


    return (
        <div>
            <TitleWithCountHeader title="Direct messages" count={dms.length} />
            <EmptyStateHandler data={dms} errorMessage={errorMessage} keyword={'direct messages'}>
                <div className="flex flex-col overflow-auto z-10 space-y-2">
                    <Tabs>
                        {groups?.map((group, index) => (
                            <TabItem key={index} text={group} onClick={() => setCurrentTab(group)} active={currentTab === group}/>
                        ))}
                    </Tabs>
                    {dms.filter((data, index) => data.groupName === currentTab).map((dm, index) => (
                       <div key={index}>
                           <MessageCard message={dm.message}
                                        timestamp={dm.timestamp} userId={dm.userId}
                                        humanReadableDate={dm.humanReadableDate}
                                        user={dm.user}/>
                       </div>
                    ))}
                </div>
            </EmptyStateHandler>
        </div>
    )
};

export default DirectMessagesContent;
