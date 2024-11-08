// assets
import { GithubOutlined, QuestionOutlined, BugOutlined } from '@ant-design/icons';
import DiscordOutlined from '../assets/images/icons/discord.svg'
import DiscordOutlinedWhite from '../assets/images/icons/discord_white.svg'
import { useTheme } from '@mui/material/styles';

// ==============================|| MENU ITEMS - SAMPLE PAGE & DOCUMENTATION ||============================== //

const DiscordOutlinedIcon = (props) => {
    const theme = useTheme();
    return (
        <img src={
            theme.palette.mode === 'dark' ? DiscordOutlinedWhite : DiscordOutlined} width="16px" alt="Discord" {...props} />
    );
};

const support = {
    id: 'support',
    title: 'menu-items.support',
    type: 'group',
    children: [
    ]
};

export default support;
