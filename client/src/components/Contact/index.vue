<template>
    <section class="section contact">
        <h2>CONTACT</h2>
        <Socials/>
        <a class="resume" href="https://drive.google.com/file/d/1VlqZZ3CyWsqqwzTgMT2CSnWSOi29lkg8/view?usp=sharing" target="__blank">RESUME</a>
    </section>
</template>

<script>
import Socials from './Socials';
export default {
    name: "Contact",
    components: {Socials},
}
</script>

<style lang="scss">
    section.contact {
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
        padding: 48px;
        width: 100vw;
        background: #0e0e0e;
        a.resume {
            text-decoration: none;
            color: white;
            font-size: 22px;
            font-weight: bold;
            letter-spacing: 4px;
            border: solid white 2px;
            border-radius: 16px;
            padding: 5px 20px;
        }
        h2 {
            letter-spacing: 14px;
            font-size: 30px;
            margin: 0;
            margin-bottom: 20px;
        }
    }

</style>
