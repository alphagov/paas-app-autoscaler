package org.cloudfoundry.autoscaler.api.util;

import org.junit.Test;

import com.sun.jersey.test.framework.JerseyTest;

import static org.junit.Assert.fail;

import org.cloudfoundry.autoscaler.api.util.CloudFoundryManager;
import org.cloudfoundry.autoscaler.api.Constants;
import org.cloudfoundry.autoscaler.api.exceptions.ClientIDLoginFailedException;

public class CloudFoundryLoginTest extends JerseyTest{
		
	public  CloudFoundryLoginTest() throws Exception{
		super("org.cloudfoundry.autoscaler.api.rest.mock.cc");
	}
	
	@Test
	public void testCloudFoundryLogin(){
		CloudFoundryManager manager = new CloudFoundryManager();
		try{
			manager.login();			
		}
		catch (Exception e){
			fail("Get exception when login");
		}
	}

	@Test(expected=ClientIDLoginFailedException.class)
	public void testCloudFoundryLoginWithNonExistClient() throws Exception{
		CloudFoundryManager manager = new CloudFoundryManager("non_exist_ID", "non_exist_secret", ConfigManager.get(Constants.CFURL) );
		manager.login();
	}

	@Test(expected=ClientIDLoginFailedException.class)
	public void testCloudFoundryLoginWithBlankClientID() throws Exception{
		CloudFoundryManager manager = new CloudFoundryManager("", ConfigManager.get(Constants.CLIENT_SECRET), ConfigManager.get(Constants.CFURL) );
		manager.login();
	}

	@Test(expected=ClientIDLoginFailedException.class)
	public void testCloudFoundryLoginWithBlankClientSecret() throws Exception{
		CloudFoundryManager manager = new CloudFoundryManager(ConfigManager.get(Constants.CLIENT_ID), "", ConfigManager.get(Constants.CFURL) );
		manager.login();
	}

	
}
