import React, {useContext, useEffect, useState} from 'react';
import EmptyStateHandler from "../../components/EmptyStateHandler/EmptyStateHandler";
import {usePublicChannels} from "./hooks/usePublicChannels";
import {GlobalContext} from "../../infrastructure/context/global.context";
import MessageCard from "../../components/Card/MessageCard";
import TitleWithCountHeader from "../../components/TitleWithCountHeader";
import TabItem from "../../components/Tabs/TabItem";
import Tabs from "../../components/Tabs/Tabs";

const PublicChannelsContent = () => {
    const {searchByUser, dateRange: [from, to]} = useContext(GlobalContext);
    const {publicChannels, groups, errorMessage} = usePublicChannels({searchByUser, from, to});
    const [currentTab, setCurrentTab] = useState(null);

    useEffect(() => {
        if (publicChannels.length > 0 && currentTab === null) {
            setCurrentTab(publicChannels[0].groupName);
        }
        return () => {
            // setCurrentTab(null);
        }
    }, [publicChannels])

    return (
        <div>
            <TitleWithCountHeader title="Public channels" count={publicChannels.length} />
            <EmptyStateHandler data={publicChannels} errorMessage={errorMessage} keyword={'public channels'}>
                <div className="flex flex-col overflow-auto z-10 space-y-2">
                    <Tabs>
                        {groups?.map((group, index) => (
                            <TabItem key={index} text={group} onClick={() => setCurrentTab(group)} active={currentTab === group}/>
                        ))}
                    </Tabs>
                    {publicChannels.filter((data, index) => data.groupName === currentTab).map((puc, index) => (
                        <div key={index}>
                            <MessageCard message={puc.message}
                                         timestamp={puc.timestamp} userId={puc.userId}
                                         humanReadableDate={puc.humanReadableDate}
                                         user={puc.user}/>
                        </div>
                    ))}
                </div>
            </EmptyStateHandler>
        </div>

    );
};

export default PublicChannelsContent;
