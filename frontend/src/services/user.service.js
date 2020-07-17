/** @format */

import { authHeader } from "../helpers/auth-header.js";
import axios from "axios";
const FileDownload = require("js-file-download");

export const userService = {
  getMyNetworks,
  getNetworkDetail,
  upNetwork,
  downNetwork,
  removeNetwork,
  addNetwork,
  generateChaincode,
  downloadServer,
  genChaincode,
  installChaincode,
  installChaincodeByConfig,
  downloadChaincode,
  upgradeChaincodeByConfig,
  uploadChaincodeFile,
  upgradeChaincode,
};

async function getMyNetworks() {
  try {
    let response = await axios.get(
      `${process.env.VUE_APP_API_BACKEND}/network/`,
      {
        headers: authHeader(),
      }
    );
    return response.data.networks;
  } catch (error) {
    throw error;
  }
}

async function getNetworkDetail(domain) {
  try {
    let response = await axios.get(
      `${process.env.VUE_APP_API_BACKEND}/network/${domain}`,
      {
        headers: authHeader(),
      }
    );
    return response.data;
  } catch (error) {
    throw error;
  }
}

async function upNetwork(domain) {
  try {
    let response = await axios.post(
      `${process.env.VUE_APP_API_BACKEND}/network/upNetwork`,
      { domain: domain },
      {
        headers: authHeader(),
      }
    );

    response = await axios.post(
      `${process.env.VUE_APP_API_BACKEND}/network/genServerWithoutChaincode`,
      { domain: domain },
      {
        headers: authHeader(),
      }
    );

    location.reload(true);
    return response;
  } catch (error) {
    throw error;
  }
}

async function downNetwork(domain) {
  try {
    let response = await axios.post(
      `${process.env.VUE_APP_API_BACKEND}/network/downNetwork`,
      { domain: domain },
      {
        headers: authHeader(),
      }
    );
    location.reload(true);
    return response;
  } catch (error) {
    throw error;
  }
}

async function removeNetwork(domain) {
  try {
    let response = await axios.post(
      `${process.env.VUE_APP_API_BACKEND}/network/removeNetwork`,
      { domain: domain },
      {
        headers: authHeader(),
      }
    );
    return response;
  } catch (error) {
    throw error;
  }
}

async function addNetwork(form) {
  try {
    let response = await axios.post(
      `${process.env.VUE_APP_API_BACKEND}/network`,
      { domain: form.domain, orgs: form.orgs },
      {
        headers: authHeader(),
      }
    );
    return response;
  } catch (error) {
    throw error;
  }
}

async function downloadServer(domain) {
  try {
    let response = await axios.post(
      `${process.env.VUE_APP_API_BACKEND}/chaincode/downloadServer`,
      { domain: domain },
      {
        headers: authHeader(),
        responseType: "arraybuffer",
      }
    );
    FileDownload(response.data, "server.zip");
    return response;
  } catch (error) {
    throw error;
  }
}

async function downloadChaincode(domain) {
  try {
    let response = await axios.post(
      `${process.env.VUE_APP_API_BACKEND}/chaincode/downloadChaincode`,
      { domain: domain },
      {
        headers: authHeader(),
        responseType: "arraybuffer",
      }
    );
    FileDownload(response.data, domain + ".go");
    return response;
  } catch (error) {
    throw error;
  }
}

async function uploadChaincodeFile(domain, file) {
  try {
    const formData = new FormData();
    formData.append("chaincodeFile", file);
    formData.append("domain", domain);

    let response = await axios.post(
      `${process.env.VUE_APP_API_BACKEND}/chaincode/uploadChaincodeFile`,
      formData,
      {
        headers: authHeader(),
      }
    );
    location.reload(true);
    return response;
  } catch (error) {
    throw error;
  }
}

async function generateChaincode(form) {
  try {
    let response = await axios.post(
      `${process.env.VUE_APP_API_BACKEND}/chaincode/upConfig`,
      { domain: form.domain, objects: form.objects },
      {
        headers: authHeader(),
      }
    );
    return response;
  } catch (error) {
    throw error;
  }
}

async function installChaincodeByConfig(domain) {
  try {
    let response = await axios.post(
      `${process.env.VUE_APP_API_BACKEND}/chaincode/installChaincode`,
      { domain: domain },
      {
        headers: authHeader(),
      }
    );

    response = await axios.post(
      `${
        process.env.VUE_APP_API_BACKEND
      }/chaincode/genServerWithChaincodeByConfig`,
      { domain: domain },
      {
        headers: authHeader(),
      }
    );
    location.reload(true);
    return response;
  } catch (error) {
    throw error;
  }
}

async function upgradeChaincodeByConfig(domain) {
  try {
    let response = await axios.post(
      `${process.env.VUE_APP_API_BACKEND}/chaincode/upgradeChaincode`,
      { domain: domain },
      {
        headers: authHeader(),
      }
    );

    response = await axios.post(
      `${
        process.env.VUE_APP_API_BACKEND
      }/chaincode/genServerWithChaincodeByConfig`,
      { domain: domain },
      {
        headers: authHeader(),
      }
    );
    location.reload(true);
    return response;
  } catch (error) {
    throw error;
  }
}

async function genChaincode(domain) {
  try {
    let response = await axios.post(
      `${process.env.VUE_APP_API_BACKEND}/chaincode/genChaincode`,
      { domain: domain },
      {
        headers: authHeader(),
      }
    );
    return response;
  } catch (error) {
    throw error;
  }
}

async function installChaincode(domain) {
  try {
    let response = await axios.post(
      `${process.env.VUE_APP_API_BACKEND}/chaincode/installChaincode`,
      { domain: domain },
      {
        headers: authHeader(),
      }
    );
    location.reload(true);
    return response;
  } catch (error) {
    throw error;
  }
}

async function upgradeChaincode(domain) {
  try {
    let response = await axios.post(
      `${process.env.VUE_APP_API_BACKEND}/chaincode/upgradeChaincode`,
      { domain: domain },
      {
        headers: authHeader(),
      }
    );
    location.reload(true);
    return response;
  } catch (error) {
    throw error;
  }
}
