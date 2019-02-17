<template>
  <div class="name-bio">
    <h2 class="name">{{first}} {{middle}} {{last}}</h2>
    <h3 class="intro">
      <StyledSVG logo="work"/>
      {{title}}
      <StyledSVG logo="location"/>
      {{location}}
    </h3>
    <p class="description">
      {{description}}
      <span>
        You can check out my recent projects on my
        <a href="https://github.com/reynld" target="__blank">Github</a> and
        <a href="https://www.linkedin.com/in/reynld/" target="__blank">LinkedIn</a>.
      </span>
    </p>
  </div>
</template>

<script>
import StyledSVG from "../../shared/SVG.js";

export default {
  name: "Description",
  components: {
    StyledSVG
  },
  data: () => ({
    first: "chris",
    middle: "reynaldo",
    last: "cedeno",
    title: "Software Engineer",
    location: "NYC",
    description:
      "Fullstack Software Engineer, team player, powerlifter, streetwear enthusiast, and devourer of japanese cuisine. I am looking for opportunities where I can impact and improve people's lives through code."
  })
};
</script>

<style lang="scss" scoped>
.name-bio {
  width: 70%;
  flex-direction: column;
  justify-content: flex-start;
  align-items: flex-start;
  position: relative;
  @media (max-width: 620px) {
    width: 100%;
    justify-content: center;
    align-items: center;
  }
  h2.name {
    font-family: "Montserrat";
    font-size: 44px;
    font-weight: 800;
    margin: 0px;
    line-height: 50px;
    user-select: none;
    text-transform: capitalize;
    @media (max-width: 555px) {
      width: 300px;
    }
  }
  h3.intro {
    font-weight: 600;
    font-size: 24px;
    line-height: 28px;
    user-select: none;
    @media (max-width: 620px) {
      line-height: 38px;
    }
  }
  p.description {
    font-size: 18px;
    width: 100%;
    a {
      color: white;
      text-decoration: underline;
    }
  }
}
</style>
