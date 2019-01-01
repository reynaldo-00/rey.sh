<template>
    <div>
    </div>
</template>

<script>
import axios from 'axios';
export default {
    name: 'Spotify',
    data: () => ({
        currentlyPlaying: [],
    }),
    methods: {
        getSong: async function() {
            const songRes = await axios('http://localhost:9001/spotify');
            this.currentlyPlaying = songRes.data;
        }
    },
    beforeMount() {
        this.getSong();
    }
}
</script>

<style>

</style>
