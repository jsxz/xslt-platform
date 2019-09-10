<?xml version="1.0" encoding="UTF-8"?>
<xsl:stylesheet xmlns:xsl="http://www.w3.org/1999/XSL/Transform"
    version="1.0">
    <xsl:template match="/">
        <xsl:for-each select="/catalog/pages"> 
            
            <first-page>
                <xsl:value-of select="normalize-space(first-page)"/>
            </first-page> 
            
            <last-pagee>
                <xsl:value-of select="normalize-space(last-pagee)"/> 
                
            </last-pagee> 
            
        </xsl:for-each>   
    </xsl:template>

    
</xsl:stylesheet>