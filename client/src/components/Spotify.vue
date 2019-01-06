<template>
    <transition name="slide">
        <a 
            v-if="currentlyPlaying" 
            :href="getSongLink" 
            target="_blank" 
            class="spotify-wrapper"
        >
            <StyledSVG logo="spotify"/>
            <span class="info-container">
                <h3>{{ getSongAndArtist }}</h3>
                <SoundBars/>
            </span>
            <img :src="getAlbumPhoto"/>
        </a>
    </transition>
</template>

<script>
import axios from 'axios';
import SoundBars from './SoundBars'
import StyledSVG from './SVG.js'

export default {
    name: 'Spotify',
    components: {
        StyledSVG,
        SoundBars
    },
    data: () => ({
        currentlyPlaying: null,
    }),
    methods: {
        getCurrentlyPlaying: async function() {
            const songRes = await axios('http://localhost:9001/spotify');
            this.currentlyPlaying = songRes.data;
        },
    },
    computed: {
        getAlbumPhoto() {
            if (!!this.currentlyPlaying && !!this.currentlyPlaying.item) {
                return this.currentlyPlaying.item.album.images[2].url
            }
            return null;
        },
        getSongAndArtist() {
            if (!!this.currentlyPlaying && !!this.currentlyPlaying.item) {
                const artist = this.currentlyPlaying.item.artists[0].name;
                const song = this.currentlyPlaying.item.name
                return `${artist} - ${song}`
            } 
            return null;
        },
        getSongLink() {
            if (!!this.currentlyPlaying && !!this.currentlyPlaying.item) {
                return this.currentlyPlaying.item.external_urls.spotify;
            }
            return '#';
        },
        animateIn() {
            if (!!this.currentlyPlaying && !!this.currentlyPlaying.item) {
                return 'slideIn'
            }
            return ''
        }
    },
    beforeMount() {
        this.getCurrentlyPlaying();
    },
}
</script>

<style scoped>
    .slide-enter-active, .slide-leave-active {
        transition: transform 550ms;
    }
    .slide-enter{
        transform: translateX(300px);
    }
    .slide-leave-to /* .fade-leave-active below version 2.1.8 */ {
        transform: translateX(-300px);
    }

    a {
        position: fixed;
        top: 0;
        right: 0px;
        width: 300px;
        height: 64px;
        background-color: #2AB759;
        color: white;
        border-bottom-left-radius: 30px;
        z-index: 1;
        display: flex;
        justify-content: space-between;
        align-items: center;
        padding-left: 10px;
        cursor: pointer;
        user-select: none;
        text-decoration: none;
    }
    img {
        margin-left: 20px;
    }

    .info-contianer {
        display: flex;
        flex-direction: column;
        justify-content: flex-start;
        align-items: center;
        height: 100%;
        padding: 10px 10px;
        width: 200px;
    }

    h3 {
        font-size: 16px;
        font-size: 1.05vw;
        font-weight: 600;
        margin: 0px;
        z-index: 99;
        text-shadow: 1px 1px 3px #000;
    }
</style>
