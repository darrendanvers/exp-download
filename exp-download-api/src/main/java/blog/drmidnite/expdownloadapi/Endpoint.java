package blog.drmidnite.expdownloadapi;

import java.io.IOException;
import javax.servlet.http.HttpServletResponse;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class Endpoint {

    // Basic test, accessible at http://localhost:8080/.
    @GetMapping
    public String get() {
        return "OK";
    }

    // Mocks up an endpoint that is generating a file for download. Accessible at http://localhost:8080/test/download.txt.
    // You can replace "test" with whatever.
    @GetMapping("/{id}/download.txt")
    public void doDownload(@PathVariable("id")String id, HttpServletResponse response) throws IOException {
        
        response.getWriter().print("You requested resource ");
        response.getWriter().print(id);
        response.getWriter().println(" from the Java server.");
    }
}
