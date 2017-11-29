
package hello;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.boot.CommandLineRunner;

import org.springframework.boot.web.client.RestTemplateBuilder;
import org.springframework.context.annotation.Bean;
import org.springframework.web.client.RestTemplate;





@SpringBootApplication
public class Application {
	
	private static final Logger log = LoggerFactory.getLogger(Application.class);
    public static void main(String[] args) {
    	System.setProperty("http.proxyHost", "proxy.cognizant.com");
    	System.setProperty("http.proxyPort", "6050");
    	System.setProperty("https.proxyHost", "proxy.cognizant.com");
        System.setProperty("https.proxyPort", "6050");
//    	Authenticator.setDefault(new Authenticator() {
//    	    protected PasswordAuthentication getPasswordAuthentication() {
//
//    	        return new PasswordAuthentication("cts\\432379","Determined*123".toCharArray());
//    	    }
//    	});
    	SpringApplication.run(Application.class);
       
        
    }
    @Bean
	public RestTemplate restTemplate(RestTemplateBuilder builder) {
		return builder.build();
	}

	@Bean
	public CommandLineRunner run(RestTemplate restTemplate) throws Exception {
		return args -> {
			Quote quote = restTemplate.getForObject(
					"http://gturnquist-quoters.cfapps.io/api/random", Quote.class);
			log.info(quote.toString());
		};
	}
}