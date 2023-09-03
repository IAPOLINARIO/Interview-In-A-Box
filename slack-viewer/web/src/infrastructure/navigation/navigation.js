import { DashboardOutlined, EmailOutlined, LockOutlined, InboxOutlined, GroupOutlined } from "@material-ui/icons";

export const navigation = [
    {
        path: '/admin/',
        label: 'Dashboard',
        icon: <DashboardOutlined />,
    },
    {
        path: '/admin/public-channels',
        label: 'Public Channels',
        icon: <EmailOutlined />,
    },
    {
        path: '/admin/private-channels',
        label: 'Private Channels',
        icon: <LockOutlined />,
    },
    {
        path: '/admin/direct-messages',
        label: 'Direct Messages',
        icon: <InboxOutlined />,
    },
    {
        path: '/admin/group-direct-messages',
        label: 'Group Direct Messages',
        icon: <GroupOutlined />,
    }
]
