import React, {useContext, useEffect, useState} from 'react';
import PrimaryContainedButton from "../../components/Buttons/PrimaryContainedButton";
import {useNavigate} from "react-router-dom";
import {InfoOutlined} from '@material-ui/icons'
import {GlobalContext} from "../../infrastructure/context/global.context";
import TitleWithCountHeader from "../../components/TitleWithCountHeader";
import EmptyStateHandler from "../../components/EmptyStateHandler/EmptyStateHandler";
import MessageCard from "../../components/Card/MessageCard";
import {useGetAllSearch} from "./hooks/useGetAllSearch";
import Tabs from "../../components/Tabs/Tabs";
import TabItem from "../../components/Tabs/TabItem";

const DashboardContent = () => {
    const {searchByUser, dateRange: [from, to]} = useContext(GlobalContext);
    const {allData, groups, errorMessage} = useGetAllSearch({searchByUser, from, to});
    const [currentTab, setCurrentTab] = useState(null);

    useEffect(() => {

        if (allData.length && currentTab === null) {
            setCurrentTab(allData[0].groupName);
        }
        return () => {
            //setCurrentTab(null);
        }
    }, [allData])

    return (
        <>
            {searchByUser === undefined || searchByUser === '' ? (
                <NoDataContent/>
            ) : (
                <div>
                    <TitleWithCountHeader title="Dashboard" count={allData.length}/>
                    <EmptyStateHandler data={allData} errorMessage={errorMessage} keyword={'public channels'}>
                        <div className="flex flex-col overflow-auto z-10 space-y-2">
                            <Tabs>
                                {groups?.map((group, index) => (
                                    <TabItem key={index} text={group} onClick={() => setCurrentTab(group)} active={currentTab === group}/>
                                ))}
                            </Tabs>
                            {allData.filter((data, index) => data.groupName === currentTab).map((puc, index) => (
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
            )}
        </>

    );
};

const NoDataContent = () => {
    const navigate = useNavigate();
    return (
        <div className="flex flex-col space-y-5 items-center justify-center max-w-4xl mx-auto">
            <div className="grid grid-cols-1 md:grid-cols-2 gap-4 w-full mt-28">
                <div className="bg-white border border-gray-200 rounded-lg shadow px-12 py-8 flex flex-col">
                    <h5 className="mb-4 text-2xl font-bold tracking-tight text-gray-900 text-center">Public
                        channels</h5>
                    {/*<h6 className="text-xl font-bold tracking-normal text-fuchsia-800 text-center mb-6">{accountingFormatNumber(334)} records</h6>*/}
                    <PrimaryContainedButton fullWidth size="lg" onClick={() => navigate('/admin/public-channels')}>See
                        more</PrimaryContainedButton>
                </div>
                <div className="bg-white border border-gray-200 rounded-lg shadow  px-12 py-8 flex flex-col">
                    <h5 className="mb-4 text-2xl font-bold tracking-tight text-gray-900 text-center">Private
                        channels</h5>
                    {/*<h6 className="text-xl font-bold tracking-normal text-fuchsia-800 text-center mb-6">{accountingFormatNumber(1552)} records</h6>*/}
                    <PrimaryContainedButton fullWidth size="lg" onClick={() => navigate('/admin/private-channels')}>See
                        more</PrimaryContainedButton>
                </div>
                <div className="bg-white border border-gray-200 rounded-lg shadow px-12 py-8 flex flex-col">
                    <h5 className="mb-4 text-2xl font-bold tracking-tight text-gray-900 text-center">Direct
                        messages</h5>
                    {/*<h6 className="text-xl font-bold tracking-normal text-fuchsia-800 text-center mb-6">{accountingFormatNumber(5562113)} records</h6>*/}
                    <PrimaryContainedButton fullWidth size="lg" onClick={() => navigate('/admin/direct-messages')}>See
                        more</PrimaryContainedButton>
                </div>
                <div className="bg-white border border-gray-200 rounded-lg shadow px-12 py-8 flex flex-col">
                    <h5 className="mb-4 text-2xl font-bold tracking-tight text-gray-900 text-center">Group direct
                        messages</h5>
                    {/*<h6 className="text-xl font-bold tracking-normal text-fuchsia-800 text-center mb-6">{accountingFormatNumber(1562113)} records</h6>*/}
                    <PrimaryContainedButton fullWidth size="lg" onClick={() => navigate('/admin/group-direct-messages')}>See
                        more</PrimaryContainedButton>
                </div>
            </div>
            <div
                className="bg-fuchsia-50/50 border border-fuchsia-800/30 text-fuchsia-900/75 px-4 py-2 w-full flex items-center space-x-2 text-lg rounded-md">
                <InfoOutlined/>
                <p>Use search bar and date range picker to set search terms and date ranges in order to filter the
                    results!</p>
            </div>
        </div>
    )
}

export default DashboardContent;
