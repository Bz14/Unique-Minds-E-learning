import img from "../app/assets/img3.jpg";

const Home = () => {
  return (
    <div
      className="relative w-full h-screen bg-cover bg-center"
      style={{ backgroundImage: `url(${img.src})` }}
    >
      <div className="absolute inset-0 bg-black opacity-50"></div>

      <div className="relative z-10 flex flex-col items-center justify-center h-full text-center text-white">
        <h1 className="text-4xl font-bold md:text-6xl">
          Welcome to Unique Minds E-learning
        </h1>
        <h2 className="text-2xl font-semibold mt-4 md:text-4xl">
          Serving the uniques !!!
        </h2>
      </div>
    </div>
  );
};

export default Home;
