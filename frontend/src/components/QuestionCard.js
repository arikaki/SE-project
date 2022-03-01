import * as React from 'react';
import Card from '@mui/material/Card';
import Fab from '@mui/material/Fab';
import CardHeader from '@mui/material/CardHeader';
// import CardMedia from '@mui/material/CardMedia';
import { red } from '@mui/material/colors';
import CardContent from '@mui/material/CardContent';
// import CardActions from '@mui/material/CardActions';
// import Collapse from '@mui/material/Collapse';
import Avatar from '@mui/material/Avatar';
import IconButton from '@mui/material/IconButton';
import EditIcon from '@mui/icons-material/Edit';
import AddIcon from '@mui/icons-material/Add';
import Button from '@mui/material/Button';
import Typography from '@mui/material/Typography';
import MoreHorizIcon from '@mui/icons-material/MoreHoriz';
import CommentIcon from '@mui/icons-material/Comment';
import { Grid } from '@mui/material';
import { makeStyles } from '@mui/styles';
import { width } from '@mui/system';

const useStyles = makeStyles((theme) => ({
    titleStyle: {
        marginLeft: '10%',
        marginTop: '1%',
        marginRight: '10%',
        // borderRadius: '15px',
        background: '#bbdefb',
        padding: '0px 10px 10px 10px',
    },
    content: {
        // flex: '1 1 auto',
        // fontsize: 10,
        overflow: 'hidden',

    },
    button: {
        height: "10px",
        width: "50px",

    }

}));

export default function QuestionScreen(props) {
    const classes = useStyles();
    return (
        <Card sx={{ maxWidth: 800 }} className={classes.titleStyle}>
            <Grid container spacing={1}>
                {<Grid item xs={9}>
                    {props.type === "answer" && <CardHeader
                        avatar={
                            <Avatar sx={{ bgcolor: red[500] }} aria-label="avatar">
                                A
                            </Avatar>


                        }
                        subheader="Alin Dobra"
                    />}
                </Grid>}
            </Grid>
            <Grid container spacing={2}>
                < Grid item xs={12} >
                    <CardContent>
                        <Typography fontSize='h6.fontSize' variant="body2" fontWeight='light' color="text.secondary" justifyContent='left' noWrap>
                            {props.title}
                        </Typography>
                        {/* <h2>{props.title}</h2> */}
                    </CardContent>
                </Grid>
                <Grid item xs={1}>
                    <Fab size='small' variant="extended" className='classes.Button'>
                        {props.type == "question" ? <EditIcon /> :
                            <CommentIcon />}
                    </Fab></Grid>
                <Grid item xs={9.5} className='classes.Button'>

                    <Fab size="small" variant="extended" className='classes.Button'>
                        <AddIcon sx={{ m: 0.1 }} />
                        <div fontSize='1px'>Follow</div>
                    </Fab>
                </Grid>
                <Grid item xs={1.5} >
                    <Button size='small' variant="contained" color="error" textsize='10px'>
                        Report
                    </Button>
                </Grid>
            </Grid>
        </Card >
    );
}
