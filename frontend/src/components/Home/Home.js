import React, { useEffect, useRef, useState } from "react";
import { ToastContainer, toast } from "react-toastify";
import { Slide } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import "./Home.css";

const Home = () => {
  const [value, setValue] = useState("identicon");
  const [pixelPattern, setPixelPattern] = useState("5");
  const [algo, setAlgo] = useState("0");
  const [identiconUrl, setIdenticonUrl] = useState("");
  const darkMode = useRef(false);

  const ServerAvailabilityMessage =
    "Please note that we are using free web hosting servers. It may take around a minute to spin up the server if it has been inactive for a few hours. Thank you for your patience.";

  const generateIdenticon = () => {
    if (value === undefined || value.trim().length == 0) {
      toast.error("Value can not be empty");
      return;
    }

    let url = `${process.env.REACT_APP_API_BASE_URL}?value=${value.trim()}&pixel_pattern=${pixelPattern}&dimension=200&algo=${algo}&dark_mode=${darkMode.current}`;

    fetch(url)
      .then((resp) => {
        if (!resp.ok) {
          return resp.text().then((errMsg) => {
            throw new Error(errMsg);
          });
        }
        return resp.blob();
      })
      .then((blob) => {
        const blobUrl = URL.createObjectURL(blob);
        setIdenticonUrl(blobUrl);
      })
      .catch((err) => {
        if (err.message === "Failed to fetch") {
          toast.error("Oops, an error occurred!");
        } else {
          toast.error(err.message.toUpperCase());
        }
      });
  };

  const handleImageError = (e) => {
    toast.error("Failed to load image.");
  };

  useEffect(() => {
    function checkServerAvailability() {
      let url = `${process.env.REACT_APP_API_BASE_URL}/health`;
      const controller = new AbortController();
      const timeoutId = setTimeout(() => controller.abort(), 2000);
      fetch(url, { signal: controller.signal })
        .then((resp) => {
          if (!resp.ok) {
            throw new Error("Some error occurred");
          }
        })
        .catch((err) => {
          if (err.name === "AbortError") {
            toast.info(ServerAvailabilityMessage, {
              autoClose: 6000,
              className: "custom-info-toast",
            });
          }
        })
        .finally(() => {
          clearTimeout(timeoutId);
        });
    }
    checkServerAvailability();
  }, []);

  return (
    <div className="App">
      <h3>Godenticon Playground</h3>
      <div className="form-group">
        <label htmlFor="valueInput">Value</label>
        <input
          type="text"
          id="valueInput"
          value={value}
          onChange={(e) => setValue(e.target.value)}
          required
        />
      </div>

      <div className="form-group">
        <label htmlFor="pixelPattern">Pixel Pattern</label>
        <select
          id="pixelPattern"
          value={pixelPattern}
          onChange={(e) => {
            setPixelPattern(e.target.value);
          }}
        >
          <option value="5">Pixel Pattern 5X5</option>
          <option value="7">Pixel Pattern 7X7</option>
          <option value="9">Pixel Pattern 9X9</option>
        </select>
      </div>

      <div className="form-group">
        <label htmlFor="algo">Algorithm</label>
        <select
          id="algorithm"
          value={algo}
          onChange={(e) => setAlgo(e.target.value)}
        >
          <option value="0">Algorithm 1</option>
          <option value="1">Algorithm 2</option>
        </select>
      </div>

      <div className="form-group">
        <label htmlFor="darkMode">Dark Mode</label>
        <label className="switch">
          <input
            type="checkbox"
            id="darkMode"
            checked={darkMode.current}
            onChange={() => {
              darkMode.current = !darkMode.current;
              generateIdenticon();
            }}
          />
          <span className="slider"></span>
        </label>
      </div>

      <div className="button-container">
        <button onClick={generateIdenticon}>Generate Identicon</button>
      </div>

      {identiconUrl && (
        <div className="image-container">
          <div
            className={`image-placeholder-${darkMode.current ? "dark" : "light"}`}
          >
            <img src={identiconUrl} onError={handleImageError} />
          </div>
        </div>
      )}
      <ToastContainer
        position="top-right"
        autoClose={2000}
        hideProgressBar
        newestOnTop={false}
        closeOnClick
        rtl={false}
        pauseOnFocusLoss={false}
        draggable
        pauseOnHover
        theme="colored"
        transition={Slide}
      />
    </div>
  );
};

export default Home;
