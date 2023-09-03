import React, {useContext, useEffect, useState} from 'react';
import EmptyStateHandler from "../../components/EmptyStateHandler/EmptyStateHandler";
import {usePrivateChannels} from "./hooks/usePrivateChannels";
import {GlobalContext} from "../../infrastructure/context/global.context";
import MessageCard from "../../components/Card/MessageCard";
import TitleWithCountHeader from "../../components/TitleWithCountHeader";
import TabItem from "../../components/Tabs/TabItem";
import Tabs from "../../components/Tabs/Tabs";

const PrivateChannelsContent = () => {
    const { searchByUser, dateRange: [from ,to] } = useContext(GlobalContext);
    const {privateChannels, groups, errorMessage} = usePrivateChannels({ searchByUser, from, to});
    const [currentTab, setCurrentTab] = useState(null);


    useEffect(() => {
        if (privateChannels.length > 0 && currentTab === null) {
            setCurrentTab(privateChannels[0].groupName);
        }
        return () => {
            // setCurrentTab(null);
        }
    }, [privateChannels])

    return (
        <div>
            <TitleWithCountHeader title="Private channels" count={privateChannels.length} />
            <EmptyStateHandler data={privateChannels} errorMessage={errorMessage} keyword={'private channels'}>
                <div className="flex flex-col overflow-auto z-10 space-y-2">
                    <Tabs>
                        {groups?.map((group, index) => (
                            <TabItem key={index} text={group} onClick={() => setCurrentTab(group)} active={currentTab === group}/>
                        ))}
                    </Tabs>
                    {privateChannels.filter((data, index) => data.groupName === currentTab).map((prc, index) => (
                        <div key={index}>
                            <MessageCard message={prc.message}
                                         timestamp={prc.timestamp} userId={prc.userId}
                                         humanReadableDate={prc.humanReadableDate}
                                         user={prc.user}/>
                        </div>
                    ))}
                </div>
            </EmptyStateHandler>
        </div>

    );
};

export default PrivateChannelsContent;
